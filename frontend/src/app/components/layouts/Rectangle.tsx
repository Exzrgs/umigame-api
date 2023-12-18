import React from 'react';

interface RectangleProps {
    color: string;
    width: number;
    height: number;
    borderRadius?: [number, number, number, number]; // 四方向のborderRadiusを持つ配列
    style?: React.CSSProperties;
}

const Rectangle: React.FC<RectangleProps> = ({ color, width, height, borderRadius = [0, 0, 0, 0], style }) => {
    const borderRadiusStyle = `${borderRadius[0]}px ${borderRadius[1]}px ${borderRadius[2]}px ${borderRadius[3]}px`;

    return (
        <div style={{
            backgroundColor: color,
            width: `${width}px`,
            height: `${height}px`,
            borderRadius: borderRadiusStyle, // 配列からスタイル文字列を生成
            ...style
        }} />
    );
}

export default Rectangle;
