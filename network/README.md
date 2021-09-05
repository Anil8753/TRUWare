# Network
Contains the network related artifacts. It uses the 'minifabric' tool that simplify the netwotk related operations.
<br/>
[minifabric github repo](https://github.com/hyperledger-labs/minifabric)
<br/>
[minifabric documentation](https://github.com/hyperledger-labs/minifabric/blob/main/docs/README.md)

Start network
If client app is outside of the containers' network
```
./minifab up -o admin.truware.com -e true
```

If client app is inside of the containers' network
```
./minifab up -o admin.truware.com
```
It brings up the network with channel name ```mychannel``` and chaincode ```siimple``` which is balance transfer between party ```a``` and party ```b```.

Bring down the network
```
./minifab down -o admin.truware.com
```

Clean everything
```
./minifab cleanup -o admin.truware.com
```

## Hyperledger Explorer

Hyperledger Explorer integration
```
./minifab explorerup 
```

```
Default username: exploreradmin
Default password: exploreradminpw
Website address:  http://192.168.1.16:7010 (192.168.1.16 is the host machine ip address)
```

Bring down Hyperledger Explorer
```
./minifab explorerdown
```

## Start up portainer web ui
While you are running your Fabric network, you can use Portainer web based management to see and interact with your running network. To start up Portainer web user interface, simply run ```./minifab portainerup``` command, to shut it down, run ```./minifab portainerdown``` command
