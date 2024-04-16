- Install the prerequisite environment
```bash
curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/blob/fea-separate-ubi/ubi/setup.sh | bash
```

- Download parameters for zk-fil:
```bash
export PARENT_PATH="<YOUR_ZK-FIL_PARAMS_PATH>"

# 32GiB parameters
curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/blob/fea-separate-ubi/ubi/fetch-param-32.sh | bash
# 512MiB parameters
curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/blob/fea-separate-ubi/ubi/fetch-param-512.sh | bash
```

- Download the binary file of `computing-provider`
```bash
curl -LO https://github.com/swanchain/go-computing-provider/releases/download/v0.4.5/computing-provider
sudo chmod +x computing-provider
```

- Create a new wallet address and recharge this address with `swan-eth`, refer to [here](https://docs.swanchain.io/swan-testnet/swan-saturn-testnet/before-you-get-started/bridge-tokens):
```bash
./computing-provider wallet new
...
0x9024a875f44172591373b92...31d67AcCEa
```
*Note: By default, set cp's repo to `~/.swan/computing`, or you can configure it by export `CP_PATH="<YOUR_CP_PATH>"`.*

- Initialize cp
```bash
./computing-provider init --ownerAddress <YOUR_WALLET_ADDRESS>
```

- Start cp service
```bash
#!/bin/bash
# Set environment variables
export FIL_PROOFS_PARAMETER_CACHE="<YOUR_ZK-FIL_PARAMS_PATH>"
export RUST_GPU_TOOLS_CUSTOM_GPU="GeForce RTX 4090:16384"   
        
nohup ./computing-provider ubi-task daemon --multi-address /ip4/<YOUR_PUBLIC_IP>/tcp/<YOUR_PORT> --node-name <YOUR_NODE_NAME> >> cp.log 2>&1 &
```
(Replace `<YOUR_ZK-FIL_PARAMS_PATH>`, `<YOUR_PUBLIC_IP>`, `<YOUR_PORT>`, and `<YOUR_NODE_NAME>` with the respective values)