import React, { useState } from 'react'

import { SignInDialog } from './SignInDialog'
import { SignUpDialog } from './SignUpDialog'

export enum AuthView {
  SignIn,
  SignUp,
}

type AuthDialogProps = {
  initialView?: AuthView
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export const AuthDialog: React.FC<AuthDialogProps> = ({
  initialView = AuthView.SignIn,
  open,
  setOpen,
}) => {
  const [authView, setAuthView] = useState<AuthView>(initialView)

  switch (authView) {
    case AuthView.SignIn:
      return (
        <SignInDialog open={open} setOpen={setOpen} setAuthView={setAuthView} />
      )
    case AuthView.SignUp:
      return (
        <SignUpDialog open={open} setOpen={setOpen} setAuthView={setAuthView} />
      )
  }
}
