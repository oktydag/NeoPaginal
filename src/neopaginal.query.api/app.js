const express = require('express')
const crawlRouter = require('./routers/crawl')

const swaggerUi = require('swagger-ui-express');
const swaggerDocument = require('./swagger.json');

const app = express()
const port = process.env.PORT || 3000

app.use(express.json())
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));
app.use(crawlRouter)

app.listen(port, () => {
    console.log('Server is up on port ' + port)
})