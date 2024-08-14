import { useNavigate } from "react-router-dom"
import { useForm, SubmitHandler, FormProvider } from "react-hook-form"
import { FormInputs } from "./types"
import { FormContainer, Heading, StyledForm, Wrapper } from "./styles"
import { SimpleInput } from "src/components"
import { useWaitForTransactionReceipt, useWriteContract } from "wagmi"
import collectionFactoryAbi from "src/abi/collectionFactory.json"
import { decodeAbiParameters } from "viem"

const CreateCollectionPage = () => {
    const methods = useForm<FormInputs>()
    const navigate = useNavigate();

    const { data: hash, writeContract, isPending } = useWriteContract()
    const { data: transactionOutput, isLoading } = useWaitForTransactionReceipt(
        { hash }
    )
    const transactionLogsData = transactionOutput?.logs?.[0]?.data

    const getDeployedCollectionAddress = (data: `0x${string}`) => {
        // find CollectionCreated event in the collection factory abi
        const eventJson = collectionFactoryAbi.find(
            (item: { name: string; type: string }) =>
                item.name === "CollectionCreated" && item.type === "event"
        )
        // return 1st element from decoded event data - deployed collection address
        return decodeAbiParameters(eventJson.inputs, data)?.[0]
    }

    const onSubmit: SubmitHandler<FormInputs> = async (data: FormInputs) => {
        writeContract({
            address: import.meta.env.VITE_FACTORY_CONTRACT_ADDRESS,
            abi: collectionFactoryAbi,
            functionName: "createCollection",
            args: [data.name, data.symbol],
        })
    }

    const goToCollection = (logsData: `0x${string}`) => {
        const address = getDeployedCollectionAddress(logsData)
        window.scrollTo(0, 0)
        navigate(`/collection/${address}`, { replace: true });
    }

    return (
        <Wrapper>
            <FormProvider {...methods}>
                <FormContainer>
                    <Heading>Create Collection</Heading>
                    <StyledForm onSubmit={methods.handleSubmit(onSubmit)}>
                        <SimpleInput
                            label="Collection name"
                            name="name"
                            options={{ required: true }}
                        />
                        <SimpleInput
                            label="Collection symbol"
                            name="symbol"
                            options={{ required: true }}
                        />
                        <input
                            type="submit"
                            value="Create Collection"
                            disabled={isLoading || isPending}
                        />
                        {isLoading ? <span>Transaction pending...</span> : null}
                        {transactionLogsData ? (
                            <>
                                <span>
                                    Success! Here's your collection
                                    address:&nbsp;
                                    {getDeployedCollectionAddress(
                                        transactionLogsData
                                    )}
                                </span>
                                <button onClick={() => goToCollection(transactionLogsData)}>
                                    Continue
                                </button>
                            </>
                        ) : null}
                    </StyledForm>
                </FormContainer>
            </FormProvider>
        </Wrapper>
    )
}

export default CreateCollectionPage
