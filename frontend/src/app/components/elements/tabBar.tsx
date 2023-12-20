import React, { ReactNode } from 'react';
import Rectangle from '../layouts/Rectangle';
import VStack from '../layouts/VStack';
import Color from '../colors/Color';
import Logo from '../../../../public/logo.png'
import './styles.css'
import ImageCrop from '@/app/utils/ImageCrop';
export default function TabBar() {
    return (
        <>
            <div style={{ position: 'fixed', width: '80px', height: '1080px', top: 0, left: 0 }}>
                <Rectangle color={Color.White} width={80} height={1080}></Rectangle>
            </div>
            <div className="my-component" style={{ position: 'fixed', zIndex: 150, width: '80px', height: '1080px', top: 0, left: 0 }} >
                <ImageCrop width='64px' height='64px' image={Logo.src}></ImageCrop>
            </div >
        </>
    )
}
