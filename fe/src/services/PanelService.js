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
  let scripts = []

  if (doc.querySelectorAll('script').length > 0) {
    const stripped = StripPanelScripts(doc)
    doc.body = stripped.body
    scripts = stripped.scripts
  }

  const body = doc.body
  const bodyContents = body.innerHTML

  const panelBody = document.createElement('div')
  panelBody.id = 'panel-html__body'
  panelBody.style = `aspect-ratio: ${aspectRatio}`
  panelBody.innerHTML = bodyContents

  if (scripts.length > 0) {
    SetPanelScripts(scripts)
  }

  return panelBody.outerHTML
}

export const StripPanelScripts = (doc) => {
  const scriptEls = doc.querySelectorAll('script')
  const scripts = []

  scriptEls.forEach((script) => {
    if (script.getAttribute('no-compile') != '') scripts.push(script.innerHTML)
    script.remove()
  })

  return { body: doc.body, scripts }
}

export const SetPanelScripts = (scripts) => {
  scripts.forEach((script) => {
    const scriptEl = document.createElement('script')
    scriptEl.setAttribute('custom_panel_script', true)
    scriptEl.innerHTML = script
    document.body.appendChild(scriptEl)
  })
}

export const RemovePanelScripts = () => {
  const scripts = document.querySelectorAll('script[custom_panel_script]')
  scripts.forEach((script) => {
    script.remove()
  })
}

export const PanelButtonListeners = (panelEl, callback) => {
  panelEl.querySelectorAll('[mcrm__button]').forEach((button) => {
    button.addEventListener('click', () => {
      callback(button)
    })
  })
}

export const PanelDialogListeners = (panelEl) => {
  panelEl.querySelectorAll('[mcrm__dialog-trigger]').forEach((dialogTrigger) => {
    const dialogEl = document.querySelector(dialogTrigger.getAttribute('dialog-trigger'))

    if (dialogEl) {
      dialogTrigger.addEventListener('click', () => {
        dialogEl.show()
      })
    }
  })

  document.querySelectorAll('dialog, dialog .dialog__close').forEach((dialogClose) => {
    dialogClose.addEventListener('click', (e) => {
      if (
        e.target.classList.contains('dialog__close') ||
        e.target.closest('.dialog__close') ||
        e.target.tagName == 'DIALOG'
      ) {
        dialogClose.closest('dialog').close()
      }
    })
  })
}
