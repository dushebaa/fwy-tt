import styled from "styled-components"

export const Wrapper = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
`

export const FormContainer = styled.div`
    width: 40%;
    display: flex;
    flex-direction: column;

    @media screen and (max-width: 767px) {
        width: 100%;
        padding: 0 2em;
        box-sizing: border-box;
    }
`

export const Heading = styled.div`
    font-size: 32px;
    font-weight: bold;
`

export const StyledForm = styled.form`
    display: flex;
    flex-direction: column;
    gap: 24px;
`