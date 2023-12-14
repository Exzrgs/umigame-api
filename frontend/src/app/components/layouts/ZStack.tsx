import React, { ReactNode } from 'react'

interface ZStackProps {
    children: ReactNode
    style?: React.CSSProperties
}

const ZStack: React.FC<ZStackProps> = ({ children, style }) => {
    return (
        <div style={{ position: 'relative', ...style }}>
            {React.Children.map(children, child => (
                <div style={{ position: 'absolute', top: 0, left: 0 }}>
                    {child}
                </div>
            ))}
        </div>
    )
}

export default ZStack
