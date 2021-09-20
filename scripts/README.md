# scripts
Contains the scripts for
- Bullet build chaincode
- Buttet build api server
- Bullet frontend web app

`build_all` - build web applications for all the nodes (warehouse and customer at present)
<br/>
`copy_cc` - copies the chaincode to network/vars/chaincode directory
<br/>
`install_cc` - install the chaincode on all the peers. It takes the chaincode name and version as parameters
<br/>
`network` - this is the main script, takes care of network down, up, reload, clean action. docker must be running before trying network script
