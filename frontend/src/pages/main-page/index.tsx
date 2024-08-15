import { useNavigate } from "react-router-dom"
import { CollectionCard, CollectionCardWrapper, Content, LowerBlock, Wrapper } from "./styles"
import useSWR from "swr"
import { CreatedCollectionResponse } from "../../common/types"
import { getReducedAddress } from "../../helpers"

const MainPage = () => {
    const navigate = useNavigate()
    const { data, isLoading } = useSWR("api/created-collections")

    const goToCollection = (address: string) => {
        navigate(`/collection/${address}`, { replace: true })
    }

    return (
        <Wrapper>
            <Content>
                <h1>hello</h1>
                <h2>pls mint you collection here</h2>
                <button onClick={() => navigate("/create")}>click</button>
            </Content>

            <LowerBlock>
                <h3>Created collections</h3>
                {isLoading ? "loading..." : null}

                {data?.length ? <CollectionCardWrapper>
                    {data.map((item: CreatedCollectionResponse, index: number) => (
                        <CollectionCard key={index}>
                            <span>Name: {item.Name}</span>
                            <span>Symbol: {item.Symbol}</span>
                            <button
                                onClick={() => goToCollection(item.Collection)}
                            >
                                Address: {getReducedAddress(item.Collection)}
                            </button>
                        </CollectionCard>
                    ))}
                </CollectionCardWrapper> : (isLoading ? null : <span>Nothing found 😔</span>)}
            </LowerBlock>
        </Wrapper>
    )
}

export default MainPage
