import { useConnect } from "wagmi"
import { ButtonContainer, ConnectorButton } from "./styles.tsx"

const ConnectButton = () => {
    const { connectors, connect } = useConnect()

    return <ButtonContainer>
        {connectors.slice(1).map((connector) => (
            <ConnectorButton
                key={connector.uid}
                onClick={() => connect({ connector })}
            >
                {connector.icon ? (
                    <img
                        alt={`${connector.name}_icon`}
                        src={connector.icon}
                        width={"16px"}
                        height={"16px"}
                    />
                ) : null}
                {connector.name}
            </ConnectorButton>
        ))}
    </ButtonContainer>
}

export default ConnectButton
