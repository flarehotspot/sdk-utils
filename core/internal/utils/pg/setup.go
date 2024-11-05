package pg

import (
	"fmt"
	"os"
	"path/filepath"
	stdstr "strings"
	"time"

	"core/internal/utils/cmd"

	gouci "github.com/digineo/go-uci"
	fs "github.com/flarehotspot/go-utils/fs"
	paths "github.com/flarehotspot/go-utils/paths"
	"github.com/goccy/go-json"
)

var (
	configPath = filepath.Join(paths.ConfigDir, "database.json")
	srvPgDir   = "/srv/pg/"
)

func SetupDb(dbpass string, dbname string) error {
	if isInstalled() {
		return nil
	}

	if err := prepPgConf(); err != nil {
		return err
	}

	if err := prepPgSrvDir(); err != nil {
		return err
	}

	if err := prepPgSrvConf(); err != nil {
		rmPgSrvDir()
		return err
	}

	if err := installPg(); err != nil {
		rmPgSrvDir()
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
		rmPgSrvDir()
		return err
	}

	return nil
}

func isInstalled() bool {
	return fs.Exists(srvPgDir) && fs.Exists(configPath)
}

func prepPgConf() error {
	// TODO: update to a correct postgresql configuration
	pgConfPath := "/etc/postgresql/my.cnf"
	bytes, err := os.ReadFile(pgConfPath)
	if err != nil {
		return err
	}

	content := string(bytes)
	if stdstr.Contains(content, "[mysqld]") {
		return nil
	}

	content += "\n"
	content += "[mysqld]\n"
	content += fmt.Sprintf("datadir = %s\n", srvPgDir)
	content += "tmpdir  = /tmp\n"

	return os.WriteFile(pgConfPath, []byte(content), 0644)
}

func prepPgSrvDir() error {
	// TODO: replace to appropriate postgresql service and directory
	commands := []string{
		"mkdir -p " + srvPgDir,
		"chown -R mariadb:mariadb " + srvPgDir,
	}

	for _, c := range commands {
		err := cmd.Exec(c, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func prepPgSrvConf() error {
	// TODO: replace to appropriate postgresql service and equivalent service commands
	values, ok := gouci.Get("mysqld", "general", "enabled")
	enabled := ok && len(values) > 0 && values[0] == "1"
	if !enabled {
		gouci.Set("mysqld", "general", "enabled", "1")
		return gouci.Commit()
	}
	return nil
}

func installPg() error {
	// TODO: replace to appropriate postgresql install commands
	commands := []string{
		"mysql_install_db --force",
		"chown -R mariadb:mariadb " + srvPgDir,
		"service mysqld start",
		"service mysqld enable",
	}

	for _, c := range commands {
		err := cmd.Exec(c, nil)
		if err != nil {
			return err
		}
	}

	// allowance time for mysql to boot first
	// sleep 3s
	time.Sleep(3 * time.Second)

	return nil

}

func rmPgSrvDir() {
	cmd.Exec("rm -rf "+srvPgDir, nil)
}

func stopDb() {
	// TODO: replace to appropriate postgresql stop command
	cmd.Exec("service mysqld stop", nil)
	cmd.Exec("service mysqld disable", nil)
}

func setRootPass(dbpass string) error {
	// TODO: replace to appropriate postgresql set password command
	command := "mysqladmin -u root password " + dbpass
	return cmd.Exec(command, nil)
}

func createDb(dbname string) error {
	// TODO: replace to appropriate postgresql create db command
	return cmd.Exec("mysqladmin create "+dbname, nil)
}

func writeConfig(dbpass string, dbname string) error {
	// TODO: replace to appropriate postgresql database configuration
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
