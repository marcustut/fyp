import React, { useEffect, useMemo, useState } from 'react'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import { Icon } from '@iconify/react'
import dayjs from 'dayjs'

import { useGetSlideQuery } from '@/generated/graphql'
import { Spinner } from '@/components'
import { ComputeClient } from '@/lib/compute'
import { COMPUTE_API_URL } from '@/lib/constants'
import { toast } from 'react-toastify'

const Slide: NextPage = () => {
  const [provisioning, setProvisioning] = useState<boolean>(false)
  const { query, replace } = useRouter()
  const id = useMemo(() => query.id as string, [query])
  const [slide] = useGetSlideQuery({ variables: { id } })
  const compute = new ComputeClient({ ServerURL: `${COMPUTE_API_URL}` })

  useEffect(() => {
    if (slide.fetching || slide.error || !slide.data) return
    ;(async () => {
      setProvisioning(true)
      try {
        const instance = await compute.runInstance({ slide_id: id })
        if (typeof instance === 'string') {
          throw new Error(instance)
        }
        replace(`http://${instance.public_dns_name}`)
      } catch (err) {
        console.error(err)
        toast.error(err as Error)
      }
      setProvisioning(false)
    })()
  }, [slide])

  return (
    <div className="relative flex h-screen flex-col items-center justify-center overflow-hidden text-indigo-500">
      {slide.fetching ? (
        <>
          <Spinner className="h-10 w-10" />
          <p className="mt-2">Fetching slide...</p>
        </>
      ) : slide.error ? (
        <>
          <Icon
            icon="heroicons-outline:exclamation-circle"
            className="h-10 w-10"
          />
          <p className="mt-2">Slide not found</p>
        </>
      ) : (
        slide.data &&
        slide.data.Slide &&
        (provisioning ? (
          <>
            <Icon icon="eos-icons:three-dots-loading" className="h-10 w-10" />
            <p className="mt-2">
              Loading{' '}
              <span className="font-medium">{slide.data.Slide.name}</span>
            </p>

            <div className="fixed bottom-4 right-4 text-xs">
              last opened on{' '}
              <span className="font-medium">
                {dayjs(slide.data.Slide.updated_at).format('h:mmA, MMM D YYYY')}
              </span>
            </div>
          </>
        ) : (
          <>
            <Icon icon="heroicons-outline:badge-check" className="h-10 w-10" />
            <p className="mt-2">
              <span className="font-medium">{slide.data.Slide.name}</span> is
              loaded, redirecting you now
            </p>
          </>
        ))
      )}
    </div>
  )
}

export default Slide
