import { useFormContext } from "react-hook-form"
import { SimpleInputProps } from "./types"
import { Wrapper } from "./styles"

const SimpleInput = ({ name, label, options }: SimpleInputProps) => {
    const {
        register,
        formState: { errors },
    } = useFormContext()

    return (
        <Wrapper>
            <label>{label}</label>
            <input
                type="text"
                {...register(name, options)}
            ></input>
            {errors[name] && <span>This field is required</span>}
        </Wrapper>
    )
}

export default SimpleInput
