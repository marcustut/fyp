import { gql } from '@urql/core'

const UserFragment = gql`
  fragment CompleteUser on User {
    id
    username
    email
    full_name
    avatar_url
    bio
    created_at
    updated_at
  }
`

const UserWithAuthFragment = gql`
  ${UserFragment}
  fragment CompleteUserWithAuth on UserWithAuth {
    access_token
    expired_at
    user {
      ...UserFragment
    }
  }
`

export const SignUp = gql`
  ${UserWithAuthFragment}
  mutation SignUp($input: CreateUserInput!) {
    SignUp(input: $input) {
      ...UserWithAuthFragment
    }
  }
`

export const SignInWithUsername = gql`
  ${UserWithAuthFragment}
  mutation SignInWithUsername($input: SignInWithUsername!) {
    SignInWithUsername(input: $input) {
      ...UserWithAuthFragment
    }
  }
`

export const SignInWithEmail = gql`
  ${UserWithAuthFragment}
  mutation SignInWithEmail($input: SignInWithEmail!) {
    SignInWithEmail(input: $input) {
      ...UserWithAuthFragment
    }
  }
`
