import React, { ReactNode } from 'react';

interface VStackProps {
    children: ReactNode;
    style?: React.CSSProperties;
}

const VStack: React.FC<VStackProps> = ({ children, style }) => {
    return (
        <div style={{ display: 'flex', flexDirection: 'column', ...style }}>
            {children}
        </div>
    );
};

export default function tabBar() {
    return (
        <>
            <div style={{ width: 80, height: 1080, left: 0, top: 0, position: 'absolute' }}>
                <div style={{ width: 80, height: 1080, left: 0, top: 0, position: 'absolute', background: 'white' }}></div>
                <div style={{ width: 40, height: 40, left: 20, top: 112, position: 'absolute' }}>
                    <div style={{ width: 40, height: 40, left: 0, top: 0, position: 'absolute', background: '#D9D9D9', borderRadius: 9999 }} />
                    <div style={{ width: 24, height: 24, left: 8, top: 6, position: 'absolute', background: 'black' }}></div>
                </div>
                <div style={{ width: 40, height: 40, left: 20, top: 304, position: 'absolute' }}>
                    <div style={{ width: 40, height: 40, left: 0, top: 0, position: 'absolute', background: '#D9D9D9', borderRadius: 9999 }} />
                    <div style={{ width: 24, height: 21.60, left: 8, top: 8, position: 'absolute', background: 'black' }}></div>
                </div>
                <div style={{ width: 40, height: 40, left: 20, top: 176, position: 'absolute' }}>
                    <div style={{ width: 40, height: 40, left: 0, top: 0, position: 'absolute', background: '#D9D9D9', borderRadius: 9999 }} />
                    <div style={{ width: 30, height: 21.60, left: 5, top: 9, position: 'absolute', background: 'black' }}></div>
                </div>
                <div style={{ width: 40, height: 40, left: 20, top: 240, position: 'absolute' }}>
                    <div style={{ width: 40, height: 40, left: 0, top: 0, position: 'absolute', background: '#D9D9D9', borderRadius: 9999 }} />
                    <div style={{ width: 30, height: 27, left: 5, top: 7, position: 'absolute', background: 'black' }}></div>
                </div>
                <img style={{ width: 64, height: 64, left: 8, top: 24, position: 'absolute', borderRadius: 512 }} src="https://via.placeholder.com/64x64" />
            </div>
            <div style={{ width: 80, height: 1080, left: 0, top: 0, position: 'absolute' }}>
                <div style={{ width: 80, height: 1080, left: 0, top: 0, position: 'absolute', background: 'white' }}></div>
                <div style={{ width: 40, height: 40, left: 20, top: 112, position: 'absolute' }}>
                    <div style={{ width: 40, height: 40, left: 0, top: 0, position: 'absolute', background: '#D9D9D9', borderRadius: 9999 }} />
                    <div style={{ width: 24, height: 24, left: 8, top: 6, position: 'absolute', background: 'black' }}></div>
                </div>
                <div style={{ width: 40, height: 40, left: 20, top: 304, position: 'absolute' }}>
                    <div style={{ width: 40, height: 40, left: 0, top: 0, position: 'absolute', background: '#D9D9D9', borderRadius: 9999 }} />
                    <div style={{ width: 24, height: 21.60, left: 8, top: 8, position: 'absolute', background: 'black' }}></div>
                </div>
                <div style={{ width: 40, height: 40, left: 20, top: 176, position: 'absolute' }}>
                    <div style={{ width: 40, height: 40, left: 0, top: 0, position: 'absolute', background: '#D9D9D9', borderRadius: 9999 }} />
                    <div style={{ width: 30, height: 21.60, left: 5, top: 9, position: 'absolute', background: 'black' }}></div>
                </div>
                <div style={{ width: 40, height: 40, left: 20, top: 240, position: 'absolute' }}>
                    <div style={{ width: 40, height: 40, left: 0, top: 0, position: 'absolute', background: '#D9D9D9', borderRadius: 9999 }} />
                    <div style={{ width: 30, height: 27, left: 5, top: 7, position: 'absolute', background: 'black' }}></div>
                </div>
                <img style={{ width: 64, height: 64, left: 8, top: 24, position: 'absolute', borderRadius: 512 }} src="https://via.placeholder.com/64x64" />
            </div>
        </>
    )
}
