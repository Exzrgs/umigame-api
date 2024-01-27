import ZStack from '../layouts/ZStack';
import Rectangle from '../layouts/Rectangle';
import Color from '../colors/Color';
import React from 'react';
import myImage from '../../../../public/bbbb.png'
import Problem from '../../models/Problem/Problem';
import ImageCrop from '../../utils/ImageCrop'
import HStack from '../layouts/HStack';
import VStack from '../layouts/VStack';
import formatDate from '@/app/utils/FormatDate';
interface ProblemCardProps {
    problem: Problem;
}

const ProblemCard: React.FC<ProblemCardProps> = ({ problem }) => {
    return (

        <div style={{ position: 'relative', width: '764px', height: '320px' }}> {/* 明示的なサイズを設定 */}
            {/* この長方形は後で画像に置換 */}
            <div style={{ position: 'absolute', zIndex: 120 }}>
                <Rectangle color={Color.White} width={764} height={256} borderRadius={[40, 40, 40, 40]} />
            </div>
            <HStack>
                <div style={{ position: 'relative', zIndex: 140, width: 280, height: 256 }}>
                    <ImageCrop width="280px" height="256px" borderRadius={[40, 0, 0, 40]} image={myImage.src} />
                </div>
                <div style={{ position: 'relative', zIndex: 150 }} >
                    <Rectangle color={Color.White} width={24} height={256} />
                </div>
                <VStack>
                    <Rectangle color={Color.White} width={0} height={24} borderRadius={[40, 40, 40, 40]} />
                    <div style={{ position: 'relative', zIndex: 150, height: 40, width: 440 }} >
                        <p style={{ width: '100%', height: '100%', color: Color.Black, fontSize: 32, fontFamily: 'Inter', fontWeight: '800', wordWrap: 'break-word' }}>{problem.title}</p>
                    </div>
                    <Rectangle color={Color.White} width={0} height={18} borderRadius={[40, 40, 40, 40]} />
                    <div style={{ position: 'relative', zIndex: 150, height: 40, width: 440 }} >
                        <p style={{ width: '100%', height: '100%', color: Color.BlackGray, fontSize: 24, fontFamily: 'Inter', fontWeight: '400', wordWrap: 'break-word' }}>{problem.author}</p>
                    </div>
                    <Rectangle color={Color.White} width={0} height={16} borderRadius={[40, 40, 40, 40]} />
                    <div style={{ position: 'relative', zIndex: 150, height: 16, width: 440 }} >
                        <p style={{ width: '100%', height: '100%', color: Color.MiddleGray, fontSize: 16, fontFamily: 'Inter', fontWeight: '400', wordWrap: 'break-word' }}>{formatDate(problem.createdAt!)}</p>
                    </div>
                    <Rectangle color={Color.White} width={0} height={18} borderRadius={[40, 40, 40, 40]} />
                    <div style={{ position: 'relative', zIndex: 150, height: 72, width: 396 }}>
                        <p style={{
                            display: '-webkit-box',
                            WebkitBoxOrient: 'vertical',
                            WebkitLineClamp: 3,
                            overflow: 'hidden',
                            textOverflow: 'ellipsis',
                            width: '100%',
                            height: '100%',
                            color: Color.WhiteGray,
                            fontSize: 16,
                            fontFamily: 'Inter',
                            fontWeight: '400',
                            wordWrap: 'break-word'
                        }}>
                            {problem.statement}
                        </p>
                    </div>


                </VStack>
            </HStack >
        </div >

    );
};

export default ProblemCard;
