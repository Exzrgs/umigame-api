import React from 'react'

interface RectangleProps {
    color: string
    width: number
    height: number
    borderRadius?: number
}

const Rectangle: React.FC<RectangleProps> = ({ color, width, height, borderRadius = 0 }) => {
    return (
        <div style={{
            backgroundColor: color,
            width: `${width}px`,
            height: `${height}px`,
            borderRadius: `${borderRadius}px`
        }} />
    )
}

export default Rectangle
