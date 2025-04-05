import chokidar from 'chokidar'
import { resolve, join, dirname, normalize, posix } from 'path'
import { fileURLToPath } from 'url'
import { spawn } from 'child_process'
import { existsSync, writeFileSync } from 'fs'
// import { posix } from 'path/posix'

// Get the directory of the current script
const __dirname = dirname(fileURLToPath(import.meta.url))
const panelsDir = join(__dirname, '../panels')

// Store active Tailwind processes
const tailwindProcesses = new Map()

// Function to stop an existing Tailwind process
const stopTailwind = (dir) => {
  if (tailwindProcesses.has(dir)) {
    // console.log(`Stopping Tailwind for ${dir}...`)
    const process = tailwindProcesses.get(dir)
    process.kill('SIGTERM') // Graceful stop

    tailwindProcesses.delete(dir)

    return true
  }

  return false
}

// Function to start Tailwind
const startTailwind = (filePath) => {
  let dir = dirname(filePath).replaceAll('\\', '/')
  dir = dir.replace(/^.*\/panels/, '../panels')
  const fileExt = filePath.split('.').pop()
  // console.log(dir)

  if (fileExt !== 'html') return

  const panelName = dir
    .split(/[\/\\]/)
    .pop()
    .replace('_', ' ')

  // Create a CSS and config file for Tailwind
  const cssFile = createTailwindStyling(dir)
  const configFile = createTailwindConfig(dir)

  const outputFile = dir + '/styles.css'

  // Restart Tailwind if it's already running
  if (!stopTailwind(dir)) {
    console.log(`Starting Tailwind for panel: ${panelName}`)
  } else {
    console.log(`Restarting Tailwind for panel: ${panelName}`)
  }

  // console.log(configFile)
  // console.log(outputFile)

  // Spawn a new Tailwind process
  const tailwindProcess = spawn(
    'npx',
    [
      'tailwindcss',
      '--output',
      outputFile,
      '--config',
      `${configFile}`,
      '--content',
      dir + '/index.html',
      '--watch',
    ],
    {
      stdio: 'inherit',
      shell: true,
    },
  )

  tailwindProcesses.set(dir, tailwindProcess)
}
const createTailwindStyling = (dir) => {
  const cssFile = dir + '/tailwind.css'
  const cssContent = `@tailwind base; @tailwind components; @tailwind utilities;`

  createFile(cssFile, cssContent)
  return cssFile
}
const createTailwindConfig = (dir) => {
  const configPath = dir + '/tailwind.config.js'

  // Create a basic tailwind.config.js with purge setup
  const configContent = `
      import path from 'path';
      import { fileURLToPath } from 'url';

      const __dirname = path.dirname(fileURLToPath(import.meta.url));

      export default {
        content: [path.resolve(__dirname, "index.html")], 
        theme: { extend: {} },
        plugins: [],
        mode: "jit",
      };
    `

  createFile(configPath, configContent)
  return configPath
}

const createFile = (filepath, content) => {
  if (existsSync(filepath)) return

  writeFileSync(filepath, content.trim())

  // console.log(`Created ${filepath}`)
}

// Watch for changes to `index.html` files
const watcher = chokidar.watch(`${panelsDir}/`, {
  persistent: true,
  ignoreInitial: false,
  ignored: '*.json, *.css, *.jpg, *.jpeg, *.png, *.webp',
})

// Start or restart Tailwind when a file is added or modified
watcher.on('change', (file) => {
  startTailwind(file)
})
