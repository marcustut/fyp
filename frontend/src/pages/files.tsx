import { CheckAuth, useAuth } from '@/lib/auth'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
const FilesPage: NextPage = () => {
  return (
    <CheckAuth>
      <div>I am File</div>
    </CheckAuth>
  )
}

export default FilesPage
