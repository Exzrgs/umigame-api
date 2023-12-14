import React, { ReactNode } from 'react'

interface ScrollableContainerProps {
    children: ReactNode
    maxHeight?: number
    maxWidth?: number
    style?: React.CSSProperties
}

const ScrollableContainer: React.FC<ScrollableContainerProps> = ({
    children,
    maxHeight,
    maxWidth,
    style
}) => {
    return (
        <div
            style={{
                overflow: 'auto',
                maxHeight: maxHeight ? `${maxHeight}px` : undefined,
                maxWidth: maxWidth ? `${maxWidth}px` : undefined,
                ...style
            }}
        >
            {children}
        </div>
    )
}

export default ScrollableContainer
