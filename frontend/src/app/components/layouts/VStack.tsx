import React, { ReactNode } from 'react'

type VStackProps = {
    children: React.ReactNode;
    style?: React.CSSProperties;
    gap?: number;
    align?: 'center' | 'left' | 'right' | 'top'; // 'top' を追加
}


const VStack: React.FC<VStackProps> = ({ children, style, gap = 0, align = 'top' }) => {
    let justifyContent = 'flex-start';
    let alignItems = 'center'; // 常に子要素を横軸に沿って中央揃えにする

    switch (align) {
        case 'center':
            justifyContent = 'center';
            break;
        case 'right':
            justifyContent = 'flex-end';
            break;
        case 'left':
            justifyContent = 'flex-start';
            break;
        // 'top' はデフォルトのため、justifyContent を 'flex-start' に設定
    }

    return (
        <div
            style={{
                display: 'flex',
                flexDirection: 'column',
                gap: `${gap}px`,
                position: 'relative',
                justifyContent,
                alignItems, // この設定により子要素が横軸に沿って中央揃えされる
                ...style
            }}
        >
            {children}
        </div>
    )
}



export default VStack
