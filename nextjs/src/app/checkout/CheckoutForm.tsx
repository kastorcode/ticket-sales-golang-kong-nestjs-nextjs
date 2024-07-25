"use client"

import { PropsWithChildren } from "react"
import { useFormState } from "react-dom"
import { checkoutAction } from "~/actions"
import { ErrorMessage } from "~/components/ErrorMessage"

// @ts-expect-error
async function getCardHash({ cardName, cardNumber, expireDate, cvv }) {
  return Math.random().toString(36).substring(7)
}

type CheckoutFormProps = PropsWithChildren & {
  className ?: string
}

export function CheckoutForm ({ children, className } : CheckoutFormProps) {

  const [state, formAction] = useFormState<{error:string},{cardHash:string,email:string }>
    (checkoutAction, { error: "" })

  async function formHandler (formData : FormData) {
    const cardHash = await getCardHash({
      cardName:   formData.get("card_name")   as string,
      cardNumber: formData.get("cc")          as string,
      expireDate: formData.get("expire_date") as string,
      cvv:        formData.get("cvv")         as string,
    })
    const email = formData.get("email") as string
    formAction({ cardHash, email })
  }

  return (
    <form action={formHandler} className={className}>
      {state.error && <ErrorMessage error={state.error} />}
      <input type="hidden" name="card_hash" />
      {children}
    </form>
  )

}