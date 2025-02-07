import React from 'react'

function IndexLayout({ children } : { children: React.ReactNode }) {
  return (
    <div className='px-4 relative'>
      {children}
    </div>
  )
}

export default IndexLayout