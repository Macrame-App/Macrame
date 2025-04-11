export const SetPanelStyle = (styleStr) => {
  const styleEl = document.createElement('style')
  styleEl.setAttribute('custom_panel_style', true)
  styleEl.innerHTML = styleStr
  document.head.appendChild(styleEl)
}

export const RemovePanelStyle = () => {
  const styleEl = document.querySelector('style[custom_panel_style]')
  if (styleEl) {
    styleEl.remove()
  }
}

export const StripPanelHTML = (html, aspectRatio) => {
  const parser = new DOMParser()
  const doc = parser.parseFromString(html, 'text/html')

  const body = doc.body
  const bodyContents = body.innerHTML

  const panelBody = document.createElement('div')
  panelBody.id = 'panel-html__body'
  panelBody.style = `aspect-ratio: ${aspectRatio}`
  panelBody.innerHTML = bodyContents

  return panelBody.outerHTML
}

export const PanelButtonListeners = (panelEl, callback) => {
  panelEl.querySelectorAll('[mcrm__button]').forEach((button) => {
    button.addEventListener('click', () => {
      callback(button)
    })
  })
}
