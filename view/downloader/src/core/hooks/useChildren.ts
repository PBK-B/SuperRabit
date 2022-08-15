import { useMemo } from 'react';

export default function useChildren(children:any) {

    const _children = useMemo(() => {
        if (!children) return [];
        return Array.isArray(children) ? children : [children]
    }, [children]);

    return _children;
}