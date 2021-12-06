
const vms = require('./balancers/client');

const client = vms.Client('http://localhost:8080');

// Scenario 1: Display available vms.
client.listOfVMS()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available vms:');
        list.forEach(console.log);
    })
    .catch((e) => {
        console.log(`Problem listing available vms: ${e.message}`);
    });

// Scenario 2: Create new channel.
client.createVMS(10, false)
    .then(() => {
        console.log('=== Scenario 2 ===');
        return client.listOfVMS()
        .then((list) => {
            console.log('=== Scenario VMSAutoInsert ===');
            console.log('Available vms:');
            list.forEach(console.log);
        })
    })
    .catch((e) => {
        console.log(`Problem creating a new channel: ${e.message}`);
    });