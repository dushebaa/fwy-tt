import { useParams } from "react-router-dom"
import { Content, StyledForm, TokenCard, TokenCardWrapper, Wrapper } from "./styles"
import { getReducedAddress } from "../../helpers"
import { FormProvider, useForm } from "react-hook-form"
import { SimpleInput } from "../../components"
import { useWaitForTransactionReceipt, useWriteContract } from "wagmi"
import collectionTokenAbi from "src/abi/collectionToken.json"
import { FormInputs } from "./types"
import useSWR from "swr"
import { MintedTokenResponse } from "../../common/types"

const CollectionDetailPage = () => {
    const { address } = useParams()
    const methods = useForm<FormInputs>()
    const { data: hash, writeContract, isPending } = useWriteContract()
    const { isLoading, data: transactionOutput } = useWaitForTransactionReceipt({ hash })

    const { data: tokensMinted, isLoading: tokensMintedLoading } = useSWR(`api/tokens-minted/${address?.toLowerCase()}`)
    console.log(tokensMinted, tokensMintedLoading)

    const onSubmit = async (data: FormInputs) => {
        writeContract({
            address,
            abi: collectionTokenAbi,
            functionName: "mint",
            args: [data.token_uri],
        })
    }

    return (
        <Wrapper>
            <Content>
                <h1>Collection {getReducedAddress(address)}</h1>
                <h2>Mint an NFT</h2>
                <FormProvider {...methods}>
                    <StyledForm onSubmit={methods.handleSubmit(onSubmit)}>
                        {/* <SimpleInput
                            label="Token Id"
                            name="token_id"
                            options={{ required: false }}
                        /> */}
                        <SimpleInput
                            label="Token URI"
                            name="token_uri"
                            options={{ required: true }}
                        />
                        <input
                            type="submit"
                            value="Create NFT"
                            disabled={isLoading || isPending}
                        />
                    </StyledForm>
                </FormProvider>
                {transactionOutput ? <span>Success!</span> : null}

                {tokensMinted?.length ? <TokenCardWrapper>
                    {tokensMinted.map((item: MintedTokenResponse, index: number) => (
                        <TokenCard key={index}>
                            <span>Recipient: {getReducedAddress(item.Recipient)}</span>
                            <span>TokenId: {item.TokenId}</span>
                            <div>TokenUri: {item.TokenUri}</div>
                        </TokenCard>
                    ))}
                </TokenCardWrapper> : <span>Nothing minted here yet.</span>}
            </Content>
        </Wrapper>
    )
}

export default CollectionDetailPage
