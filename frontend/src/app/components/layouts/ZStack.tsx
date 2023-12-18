import React, { ReactNode } from 'react';

interface ZStackProps {
    children: ReactNode;
    style?: React.CSSProperties;
    gap?: number; // gapはZStackにはあまり意味がないかもしれませんが、オプションとして残します。
}

const ZStack: React.FC<ZStackProps> = ({ children, style, gap = 0 }) => {
    // 子要素に適用するスタイルを生成します。
    const childStyle: React.CSSProperties = {
        position: 'absolute',
        top: gap,
        left: gap
    };

    return (
        <div style={{ position: 'relative', ...style }}>
            {React.Children.map(children, child => (
                <div style={childStyle}>{child}</div>
            ))}
        </div>
    );
};

export default ZStack;
