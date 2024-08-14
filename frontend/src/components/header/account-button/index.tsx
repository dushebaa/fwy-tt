import { useAccount, useDisconnect } from "wagmi"
import { Wrapper } from "./styles"
import { getReducedAddress } from "../../../helpers"

const AccountButton = () => {
    const { disconnect } = useDisconnect()
    const { address, chainId } = useAccount()

    return (
        <Wrapper>
            <button disabled>{getReducedAddress(address)}@{chainId}</button>
            <button onClick={() => disconnect()}>Disconnect</button>
        </Wrapper>
    )
}

export default AccountButton
