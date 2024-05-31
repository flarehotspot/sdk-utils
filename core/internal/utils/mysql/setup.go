//go:build !dev

package mysql

import (
	"fmt"
	"os"
    "encoding/json"
	"path/filepath"
	stdstr "strings"
	"time"

	"core/internal/utils/cmd"
	"sdk/libs/go-uci"
fs	"sdk/utils/fs"
paths	"sdk/utils/paths"
)

var (
	configPath  = filepath.Join(paths.ConfigDir, "database.json")
	srvMysqlDir = "/srv/mysql"
)

func SetupDb(dbpass string, dbname string) error {
	if isInstalled() {
		return nil
	}

	if err := prepareMycnf(); err != nil {
		return err
	}

	if err := prepareSrvMysqlDir(); err != nil {
		return err
	}

	if err := prepareConfigMysqld(); err != nil {
		rmSrvMysqlDir()
		return err
	}

	if err := mysqlInstall(); err != nil {
		rmSrvMysqlDir()
		stopDb()
		return err
	}

	if err := setRootPass(dbpass); err != nil {
		return err
	}

	if err := createDb(dbname); err != nil {
		return err
	}

	if err := writeConfig(dbpass, dbname); err != nil {
		rmSrvMysqlDir()
		return err
	}

	return nil
}

func isInstalled() bool {
	return fs.Exists(srvMysqlDir) && fs.Exists(configPath)
}

func prepareMycnf() error {
	mycnf := "/etc/mysql/my.cnf"
	bytes, err := os.ReadFile(mycnf)
	if err != nil {
		return err
	}

	content := string(bytes)
	if stdstr.Contains(content, "[mysqld]") {
		return nil
	}

	content += "\n"
	content += "[mysqld]\n"
	content += fmt.Sprintf("datadir = %s\n", srvMysqlDir)
	content += "tmpdir  = /tmp\n"

	return os.WriteFile(mycnf, []byte(content), 0644)
}

func prepareSrvMysqlDir() error {
	commands := []string{
		"mkdir -p " + srvMysqlDir,
		"chown -R mariadb:mariadb " + srvMysqlDir,
	}

	for _, c := range commands {
		err := cmd.Exec(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func prepareConfigMysqld() error {
	values, ok := uci.Get("mysqld", "general", "enabled")
	enabled := ok && len(values) > 0 && values[0] == "1"
	if !enabled {
		uci.Set("mysqld", "general", "enabled", "1")
		return uci.Commit()
	}
	return nil
}

func mysqlInstall() error {
	commands := []string{
		"mysql_install_db --force",
		"chown -R mariadb:mariadb " + srvMysqlDir,
		"service mysqld start",
		"service mysqld enable",
	}

	for _, c := range commands {
		err := cmd.Exec(c)
		if err != nil {
			return err
		}
	}

	// allowance time for mysql to boot first
	// sleep 3s
	time.Sleep(3 * time.Second)

	return nil
}

func setRootPass(dbpass string) error {
	command := "mysqladmin -u root password " + dbpass
	return cmd.ExecAsh(command)
}

func createDb(dbname string) error {
	return cmd.ExecAsh("mysqladmin create " + dbname)
}

func writeConfig(dbpass string, dbname string) error {
	cfg := map[string]string{
		"host":     "localhost",
		"username": "root",
		"password": dbpass,
		"database": dbname,
	}

	bytes, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, bytes, 6004)
}

func rmSrvMysqlDir() {
	cmd.Exec("rm -rf " + srvMysqlDir)
}

func stopDb() {
	cmd.Exec("service mysqld stop")
	cmd.Exec("service mysqld disable")
}
