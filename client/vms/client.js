const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listOfVMS: () => client.get('/balancers'),
        createVMS: (id, status) => client.post('/balancers', { id, status })
    }

};

module.exports = { Client };