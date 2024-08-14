const getReducedAddress = (address?: string) => {
    if (!address) return ""
    return `${address.slice(0, 6)}...${address.slice(-5)}`
}

export default getReducedAddress
