import React from 'react';

type CircleProps = {
    size: number; // 円の直径（ピクセル単位）
    color: string; // CSSで有効な色の指定
};

const Circle: React.FC<CircleProps> = ({ size, color }) => {
    const style = {
        width: `${size}px`,
        height: `${size}px`,
        backgroundColor: color,
        borderRadius: '50%', // 円形にするために必要
        display: 'inline-block' // インラインブロック要素として表示
    };

    return <div style={style} />;
};

export default Circle;
