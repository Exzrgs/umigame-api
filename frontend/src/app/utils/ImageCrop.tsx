interface Props {
    width: string;  // 例: '200px', '100%'
    height: string; // 例: '200px', '50%'
    borderRadius?: [number, number, number, number]; // 四方向のborderRadiusを持つ配列
    image: string;  // 画像のパス
}

const ImageCrop: React.FC<Props> = ({ width, height, borderRadius = [0, 0, 0, 0], image }) => {
    const borderRadiusStyle = `${borderRadius[0]}px ${borderRadius[1]}px ${borderRadius[2]}px ${borderRadius[3]}px`;

    return (
        <div style={{
            width: width,
            height: height,
            overflow: 'hidden',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            borderRadius: borderRadiusStyle // 四方向のborderRadiusを適用
        }}>
            <img src={image} alt="クロップされた画像" style={{ minWidth: '100%', minHeight: '100%' }} />
        </div>
    );
};

export default ImageCrop
