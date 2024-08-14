import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import App from "./App.tsx"
import "./index.css"
import { WagmiProvider } from "wagmi"
import { wagmiConfig } from "./wagmi/index.tsx"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { SWRConfig } from "swr"

const queryClient = new QueryClient()

const fetcher = async (url: string, init: RequestInit) => {
    const baseUrl = import.meta.env.VITE_BACKEND_BASE_URL
    const fullUrl = `${baseUrl}${url}`
    const res = await fetch(fullUrl, init)
    return res.json()
}

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <QueryClientProvider client={queryClient}>
            <WagmiProvider config={wagmiConfig}>
                <SWRConfig value={{ fetcher }}>
                    <App />
                </SWRConfig>
            </WagmiProvider>
        </QueryClientProvider>
    </StrictMode>
)
