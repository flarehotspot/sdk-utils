#!/usr/bin/env node

const path = require('path');
const fs = require('fs-extra');
const { Octokit } = require('@octokit/core');
const execAsync = require('./exec-async.js');
const coreVersion = require('./core-version.js');
const searchFiles = require('./search-files.js');
const GITHUB_TOKEN = process.env.GITHUB_TOKEN;
const OWNER = 'flarehotspot';
const REPO = 'sdk';
const OUT_DIR = path.join(__dirname, '../.cache/devkit/output');
const octokit = new Octokit({ auth: GITHUB_TOKEN });

const main = async () => {
  const CORE_VERSION = await coreVersion();

  async function isPreRelease() {
    const preKeywords = ['alpha', 'beta', 'rc', 'pre'];
    const tag = CORE_VERSION.toLowerCase();
    for (const keyword of preKeywords) {
      if (tag.includes(keyword)) {
        return true;
      }
    }
    return false;
  }

  async function releaseNotes() {
    const notes = await fs.readFile(
      path.join(__dirname, './RELEASE_NOTES.md'),
      'utf8'
    );
    return (
      notes +
      `
---
**Download Instruction:**

If you are using Windows, Mac or Linux on x86, select the \`amd64\` zip file.
If you are using Mac on Apple silicon, select the \`arm64\` zip file.
Otherwise, select the version that's compatible with your device.
      `
    );
  }

  const { data } = await octokit.request(
    'POST /repos/{owner}/{repo}/releases',
    {
      owner: OWNER,
      repo: REPO,
      tag_name: CORE_VERSION,
      name: CORE_VERSION,
      body: await releaseNotes(),
      draft: false,
      prerelease: await isPreRelease(),
      generate_release_notes: false,
      headers: {
        'X-GitHub-Api-Version': '2022-11-28'
      }
    }
  );

  async function deleteRelease() {
    await octokit.request(
      'DELETE /repos/{owner}/{repo}/releases/{release_id}',
      {
        owner: OWNER,
        repo: REPO,
        release_id: data.id,
        headers: {
          'X-GitHub-Api-Version': '2022-11-28'
        }
      }
    );
    console.log(`Deleted release: ${data.id}`);
  }

  async function uploadZipFile(filePath) {
    await fs.ensureDir(OUT_DIR);
    const dest = path.join(OUT_DIR, path.basename(filePath));
    await fs.move(filePath, dest);
    // const fileData = await fs.readFile(filePath);
    // console.log(`Uploading ${fileData}`);
    // await octokit.request(`POST ${data.upload_url}`, {
    //   owner: OWNER,
    //   repo: REPO,
    //   name: path.basename(filePath),
    //   release_id: data.id,
    //   data: fileData,
    //   headers: {
    //     'X-GitHub-Api-Version': '2022-11-28',
    //     'Content-Type': 'application/zip'
    //   }
    // });
    // console.log(`Success uploading file: ${filePath}`);
  }

  async function zipAndUploadDevkit() {
    const releaseDirs = await searchFiles(
      path.join(__dirname, '../devkit-release'),
      async (dir, f) => {
        if (f === 'go.work') {
          const dockerFile = path.join(dir, 'Dockerfile');
          return await fs.exists(dockerFile);
        }
      },
      async (dir) => dir,
      { stopRecurse: true }
    );

    for (const dir of releaseDirs) {
      const zipPath = `${dir}.zip`;
      console.log(`Zipping ${dir} -> ${zipPath}`);
      await execAsync(`zip -r ${zipPath} .`, {
        cwd: dir
      });
      await uploadZipFile(zipPath);
      console.log(`Success uploading file: ${zipPath}`);
    }
  }

  try {
    await zipAndUploadDevkit();
  } catch (e) {
    console.log(e);
    await deleteRelease();
    process.exit(1);
  }
};

module.exports = main();
