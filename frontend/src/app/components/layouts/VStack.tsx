import React, { ReactNode } from 'react'

interface VStackProps {
    children: ReactNode
    style?: React.CSSProperties
    gap?: number
}

const VStack: React.FC<VStackProps> = ({ children, style, gap = 0 }) => {
    return (
        <div style={{ display: 'flex', flexDirection: 'column', gap: `${gap}px`, ...style }}>
            {children}
        </div>
    )
}

export default VStack
