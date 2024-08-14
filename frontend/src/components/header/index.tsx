import { useAccount } from "wagmi"
import { Wrapper } from "./styles"
import ConnectButton from "./connect-button"
import AccountButton from "./account-button"
import { useNavigate } from "react-router-dom"

const Header = () => {
    const { address } = useAccount()
    const navigate = useNavigate()

    const goToHomepage = () => {
        navigate("/", { replace: true })
    }

    return (
        <Wrapper>
            <button onClick={goToHomepage}>home</button>
            {address ? <AccountButton /> : <ConnectButton />}
        </Wrapper>
    )
}

export default Header
