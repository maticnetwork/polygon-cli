gen:
	FOUNDRY_PROFILE=lite forge build --contracts src/loadtester
	$(call gen_go_binding,LoadTester,src/loadtester,loadtester,src/loadtester)

	forge build --skip LoadTester
	$(call gen_go_binding,ERC20,src/tokens/ERC20,tokens,src/tokens)
	$(call gen_go_binding,ERC721,src/tokens/ERC721,tokens,src/tokens)
	$(call gen_go_binding,Swapper,src/tokens/Swapper,tokens,src/tokens)
	$(call gen_go_binding,WETH9,src/tokens/WETH9,tokens,src/tokens)

test:
	$(call gen_go_binding_from_lib,openzeppelin-contracts-v3.4.1-solc-0.7-2,ProxyAdmin,src/openzeppelin/ProxyAdmin,openzeppelin,src/openzeppelin)
	$(call gen_go_binding_from_lib,openzeppelin-contracts-v3.4.1-solc-0.7-2,TransparentUpgradeableProxy,src/openzeppelin/TransparentUpgradeableProxy,openzeppelin,src/openzeppelin)

# Generate simple go bindings.
# - $1: smart contract name.
# - $2: smart contract folder path.
# - $3: go package name.
# - $4: go file output path.
define gen_go_binding
	echo "Generating ABI, bytecode and go binding for $1..."
	cat ./out/$1.sol/$1.json | jq -r '.abi' 						> ./$2/$1.abi
	cat ./out/$1.sol/$1.json | jq -r '.bytecode.object' > ./$2/$1.bin
	abigen --abi ./$2/$1.abi --bin ./$2/$1.bin --pkg $3 --type $1 --out ./$4/$1.go
endef

# Generate go bindings for smart contracts under `lib/`.
# - $1: library name.
# - $2: smart contract name.
# - $3: smart contract folder path.
# - $4: go package name.
# - $5: go file output path.
define gen_go_binding_from_lib
	echo "Generating ABI, bytecode and go binding for $2..."
	cd "./lib/$1" && forge build
	cat "./lib/$1/out/$2.sol/$2.json" | jq -r '.abi' 						 > "./$3/$2.abi"
	cat "./lib/$1/out/$2.sol/$2.json" | jq -r '.bytecode.object' > "./$3/$2.bin"
	abigen --abi "./$3/$2.abi" --bin "./$3/$2.bin" --pkg $4 --type $2 --out ./$5/$2.go
endef