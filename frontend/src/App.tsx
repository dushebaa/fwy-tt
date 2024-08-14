import { BrowserRouter as Router, Routes, Route } from "react-router-dom"
import "./App.css"
import Layout from "./components/layout"
import { CollectionDetailPage, CreateCollectionPage, MainPage } from "./pages"
import {} from "react-router-dom"

function App() {
    return (
        <Router>
            <Routes>
                <Route element={<Layout />}>
                    <Route index element={<MainPage />} />
                    <Route path="/create" element={<CreateCollectionPage />} />
                    <Route
                        path="/collection/:address"
                        element={<CollectionDetailPage />}
                    />
                </Route>
            </Routes>
        </Router>
    )
}

export default App
