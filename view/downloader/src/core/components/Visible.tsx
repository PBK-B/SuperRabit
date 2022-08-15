import React, { useMemo } from 'react';

const Visible = React.memo((props:{
    children?:any;
    show?:boolean;
}) => {
    const show = useMemo(() => props.show ?? false, [props.show])
    if(show === false){
        return null;
    }
    return props.children
});
export default Visible;