import { Outlet } from "react-router-dom"
import Header from "../header"
import { Wrapper } from "./styles"

const Layout = () => {
    return <Wrapper>
        <Header />
        <Outlet />
    </Wrapper>
}

export default Layout
