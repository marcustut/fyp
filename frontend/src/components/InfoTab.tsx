import React from 'react'

type InfoTabProps = {
  mainRender: () => JSX.Element
  subRender?: () => JSX.Element
  className?: string
}

export const InfoTab: React.FC<InfoTabProps> = ({
  mainRender,
  subRender,
  className = '',
}) => {
  return (
    <div className={`w-1/5 ${className}`}>
      {mainRender()}
      {subRender && (
        <>
          <div className="my-8 h-[1px] w-full bg-slate-200" />
          {subRender()}
        </>
      )}
    </div>
  )
}
