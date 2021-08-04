const config = require('../config.json')

const authentication = async (req, res, next) => {
    try {
        const auth = {login: config.ApplicationSettings.Authentication.Username, password: config.ApplicationSettings.Authentication.Password}

        const b64auth = (req.headers.authorization || '').split(' ')[1] || ''
        const [login, password] = Buffer.from(b64auth, 'base64').toString().split(':')

        if (login && password && login === auth.login && password === auth.password) {
            return next()
        }

        res.set('WWW-Authenticate', 'Basic realm="401"')
        res.status(401).send('Authentication required.')

    } catch (e) {
        res.status(401).send({ error: 'Please authenticate.' })
    }
}

module.exports = authentication