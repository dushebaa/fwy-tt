import { createConfig, http } from "wagmi"
import { sepolia } from "wagmi/chains"
import { injected, walletConnect } from "wagmi/connectors"

const projectId = import.meta.env.VITE_WALLETCONNECT_PROJECT_ID || ""

const config = createConfig({
    chains: [sepolia],
    connectors: [injected(), walletConnect({ projectId })],
    transports: {
        [sepolia.id]: http("https://gateway.tenderly.co/public/sepolia"),
    },
})

export default config
