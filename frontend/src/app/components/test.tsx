import ProblemList from '../problems/page';
import VStack from './layouts/VStack';
import HStack from './layouts/HStack';
import { prototype } from 'events';
import ZStack from './layouts/ZStack';
import Rectangle from './layouts/Rectangle';
import Color from './colors/Color'
import ScrollableContainer from './layouts/ScrollableContainer';
import Link from 'next/link'
import ProblemCardList from './lists/ProblemCardList';
import Problem from '@/app/models/Problem';

const problems: Problem[] = [
    {
        id: 1, title: "カメオカメオカメオカメオ", statement: "カメオはカメオに殺された。そのカメオもまた、カメオに殺されてしまった。カメオはカメオになった。カメオは全てのカメオを殺した。どういう状況？",
        answer: "aaaa",
        author: "東雲篠葉",
        reference: "aaaa",
        referenceURL: "aaaa",
        createdAt: new Date()
    },
    {
        id: 2, title: "aaaa", statement: "aaaa",
        answer: "aaaa",
        author: "aaaa",
        reference: "aaaa",
        referenceURL: "aaaa",
        createdAt: new Date()
    },
    {
        id: 3, title: "aaaa", statement: "aaaa",
        answer: "aaaa",
        author: "aaaa",
        reference: "aaaa",
        referenceURL: "aaaa",
        createdAt: new Date()
    },
    {
        id: 4, title: "aaaa", statement: "aaaa",
        answer: "aaaa",
        author: "aaaa",
        reference: "aaaa",
        referenceURL: "aaaa",
        createdAt: new Date()
    }
]
export default function test() {
    return (
        <>
            <HStack style={{ height: '1080px', alignItems: 'flex-start' }}>
                <Rectangle color={Color.White} width={80} height={1080}></Rectangle>
                <Rectangle color={Color.Gray} width={80} height={1080}></Rectangle>
                <ScrollableContainer maxHeight={1080} maxWidth={1840} style={{ width: '100%' }}>
                    <VStack>
                        <Rectangle color={Color.Gray} width={1920} height={64}></Rectangle>
                        <HStack gap={80}>
                            <VStack gap={64}>
                                <ProblemCardList problems={problems}></ProblemCardList>
                            </VStack>
                            <VStack gap={64}>
                                <ProblemCardList problems={problems}></ProblemCardList>
                            </VStack>

                        </HStack>
                    </VStack >
                </ScrollableContainer >
            </HStack >
        </>
    )
}
