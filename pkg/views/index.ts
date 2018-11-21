import * as next from 'next'
import * as express from 'express'
const adminNEXTConfig = require("./next.config")



async function start() {
  const nextInstance = next({
    conf: {
      ...adminNEXTConfig,
      publicRuntimeConfig: {
        // fill any Client runtime config
        // here, and it will appear on __NEXT_DATA__
      },
    },
    dev: true,
  })

  await nextInstance.prepare()

  const requestHandler = nextInstance.getRequestHandler()
  
  const server = express()


  server.use((req, res) => {
    requestHandler(req, res)
  })

  server.listen(3001, () => {
    console.log('views ready to serve, :3001')
  })
}


start()