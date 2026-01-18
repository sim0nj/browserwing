#!/usr/bin/env node

const https = require('https');
const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const REPO = 'browserwing/browserwing';
const PACKAGE_VERSION = require('./package.json').version;

// Platform mapping
const PLATFORM_MAP = {
  darwin: 'darwin',
  linux: 'linux',
  win32: 'windows'
};

const ARCH_MAP = {
  x64: 'amd64',
  arm64: 'arm64',
  arm: 'armv7'
};

function getPlatform() {
  const platform = PLATFORM_MAP[process.platform];
  const arch = ARCH_MAP[process.arch];
  
  if (!platform || !arch) {
    throw new Error(`Unsupported platform: ${process.platform}-${process.arch}`);
  }
  
  return { platform, arch };
}

function getDownloadURL(version, platform, arch) {
  const binaryName = platform === 'windows' 
    ? `browserwing-${platform}-${arch}.exe`
    : `browserwing-${platform}-${arch}`;
  
  return `https://github.com/${REPO}/releases/download/v${version}/${binaryName}`;
}

function downloadFile(url, dest) {
  return new Promise((resolve, reject) => {
    console.log(`Downloading BrowserWing from ${url}...`);
    
    const file = fs.createWriteStream(dest);
    
    https.get(url, { 
      headers: { 'User-Agent': 'browserwing-npm-installer' }
    }, (response) => {
      // Handle redirects
      if (response.statusCode === 302 || response.statusCode === 301) {
        return https.get(response.headers.location, (redirectResponse) => {
          redirectResponse.pipe(file);
          file.on('finish', () => {
            file.close();
            resolve();
          });
        }).on('error', reject);
      }
      
      if (response.statusCode !== 200) {
        reject(new Error(`Failed to download: ${response.statusCode}`));
        return;
      }
      
      response.pipe(file);
      
      file.on('finish', () => {
        file.close();
        resolve();
      });
    }).on('error', (err) => {
      fs.unlink(dest, () => {});
      reject(err);
    });
  });
}

async function install() {
  try {
    const { platform, arch } = getPlatform();
    console.log(`Installing BrowserWing for ${platform}-${arch}...`);
    
    // Create bin directory
    const binDir = path.join(__dirname, 'bin');
    if (!fs.existsSync(binDir)) {
      fs.mkdirSync(binDir, { recursive: true });
    }
    
    // Determine binary name
    const binaryName = platform === 'windows' ? 'browserwing.exe' : 'browserwing';
    const binaryPath = path.join(binDir, binaryName);
    
    // Download binary
    const downloadURL = getDownloadURL(PACKAGE_VERSION, platform, arch);
    await downloadFile(downloadURL, binaryPath);
    
    // Make executable on Unix-like systems
    if (platform !== 'windows') {
      fs.chmodSync(binaryPath, 0o755);
    }
    
    console.log('BrowserWing installed successfully!');
    console.log('');
    console.log('Quick start:');
    console.log('  browserwing --port 8080');
    console.log('  Open http://localhost:8080');
    
  } catch (error) {
    console.error('Installation failed:', error.message);
    console.error('');
    console.error('You can manually download from:');
    console.error(`https://github.com/${REPO}/releases`);
    process.exit(1);
  }
}

install();
