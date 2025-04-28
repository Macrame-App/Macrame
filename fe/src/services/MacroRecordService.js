const keyMap = {
  // Modifier keys
  Control: 'Ctrl',
  Shift: 'Shift',
  Alt: 'Alt',
  Meta: 'Win',
  CapsLock: 'Caps',
  // Special keys
  PageUp: 'PgUp',
  PageDown: 'PgDn',
  ScrollLock: 'Scr Lk',
  Insert: 'Ins',
  Delete: 'Del',
  Escape: 'Esc',
  Space: 'Space',
  // Symbol keys
  Backquote: '`',
  Backslash: '\\',
  BracketLeft: '[',
  BracketRight: ']',
  Comma: ',',
  Equal: '=',
  Minus: '-',
  Period: '.',
  Quote: "'",
  Semicolon: ';',
  Slash: '/',
  // Arrow keys
  ArrowUp: '&#9650;',
  ArrowRight: '&#9654;',
  ArrowDown: '&#9660;',
  ArrowLeft: '&#9664;',
  // Media keys
  MediaPlayPause: 'Play',
  MediaStop: 'Stop',
  MediaTrackNext: 'Next',
  MediaTrackPrevious: 'Prev',
  MediaVolumeDown: 'Down',
  MediaVolumeUp: 'Up',
  AudioVolumeMute: 'Mute',
  AudioVolumeDown: 'Down',
  AudioVolumeUp: 'Up',
}

/**
 * Filters a keyboard event and returns an object with two properties:
 * loc (optional) and str.
 * loc is the location of the key (either 'left', 'right', or 'num').
 * str is the string representation of the key (e.g. 'a', 'A', 'Enter', etc.).
 * If the key is a modifier key, it is represented by its name (e.g. 'Ctrl', 'Shift', etc.).
 * If the key is not a modifier key, it is represented by its character (e.g. 'a', 'A', etc.).
 * If the key is not a character key, it is represented by its symbol (e.g. ',', '.', etc.).
 * @param {KeyboardEvent} e - The keyboard event to filter.
 * @return {Object} An object with two properties: loc (optional) and str.
 */
export const filterKey = (e) => {
  const k = {} // Object k (key)

  // If location is set, set loc (location)
  if (e.location === 1) k.loc = 'left'
  if (e.location === 2) k.loc = 'right'
  if (e.location === 3) k.loc = 'num'

  if (e.key.includes('Media') || e.key.includes('Audio')) k.loc = mediaPrefix(e)

  // If code is in keyMap, set str by code
  if (keyMap[e.code] || keyMap[e.key]) {
    k.str = keyMap[e.code] || keyMap[e.key]
  } else {
    // If code is not in keyMap, set str by e.key
    k.str = e.key.toLowerCase()
  }

  // return k object
  return k
}

/**
 * Returns a string prefix for the given media key.
 * @param {KeyboardEvent} e - The keyboard event to get the prefix for.
 * @return {string} The prefix for the key (either 'Media' or 'Volume').
 */
const mediaPrefix = (e) => {
  switch (e.key) {
    case 'MediaPlayPause':
    case 'MediaStop':
    case 'MediaTrackNext':
    case 'MediaTrackPrevious':
      return 'Media'
    case 'MediaVolumeDown':
    case 'MediaVolumeUp':
    case 'AudioVolumeDown':
    case 'AudioVolumeUp':
    case 'AudioVolumeMute':
      return 'Volume'
  }
}

export const isRepeat = (lastStep, e, direction) => {
  return (
    lastStep &&
    lastStep.type === 'key' &&
    lastStep.code === e.code &&
    lastStep.direction === direction
  )
}

export const invalidMacro = (steps) => {
  const downKeys = []
  const upKeys = []

  Object.keys(steps).forEach((stepKey) => {
    const step = steps[stepKey]

    if (step.type !== 'key') return

    if (step.direction == 'down') downKeys.push(step.key)
    if (step.direction == 'up') {
      if (!downKeys.includes(step.key)) upKeys.push(step.key)
      else downKeys.splice(downKeys.indexOf(step.key), 1)
    }
  })

  if (upKeys.length === 0 && downKeys.length === 0) return false

  return { down: downKeys, up: upKeys }
}

export const translateJSON = (json) => {
  const steps = []

  json.forEach((step) => {
    if (step.type === 'delay') steps.push(step)
    if (step.type === 'key') steps.push(codeToStep(step.code, step.direction))
  })

  return steps
}

export const codeToStep = (code, direction) => {
  let key = ''
  let location = 0
  let codeStr = code

  if (code.includes('Left')) {
    key = code.replace('Left', '')
    location = 1
  }
  if (code.includes('Right')) {
    key = code.replace('Right', '')
    location = 2
  }
  if (code.includes('Numpad')) {
    key = code.replace('Numpad', '')
    location = 3
  }

  if (code.includes('Media')) codeStr = ''

  if (key === '') key = code

  const stepObj = {
    type: 'key',
    code: codeStr,
    key: key,
    location: location,
    direction: direction,
  }

  return { ...stepObj, keyObj: filterKey(stepObj) }
}
