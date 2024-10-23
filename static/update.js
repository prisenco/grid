async function updateCell (room, x, y, color) {
  const url = `/room/${room}`
  const params = new URLSearchParams()
  params.set('x', x)
  params.set('y', y)
  params.set('color', color)

  try {
    const response = await fetch(url, {
      body: params,
      method: 'POST'
    })
    if (!response.ok) {
      console.log(response.status)
      return
    }
    const newBody = await response.text()
    document.body.innerHTML = newBody
  } catch (e) {
    console.error(e.message)
  }
}
