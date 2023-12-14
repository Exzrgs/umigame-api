import React, { ReactNode } from 'react'

interface HStackProps {
    children: ReactNode
    style?: React.CSSProperties
    gap?: number
}

const HStack: React.FC<HStackProps> = ({ children, style, gap = 0 }) => {
    return (
        <div style={{ display: 'flex', flexDirection: 'row', gap: `${gap}px`, ...style }}>
            {children}
        </div>
    )
}

export default HStack
