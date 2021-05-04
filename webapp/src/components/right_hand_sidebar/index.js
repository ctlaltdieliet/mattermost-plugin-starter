import React from 'react'

export default class RightSideBar extends React.PureComponent {
  
    render() {
        return <div dangerouslySetInnerHTML={{ __html: "<iframe src='https://www.youtube-nocookie.com/embed/IxAKFlpdcfc?&playlist=IxAKFlpdcfc&controls=1&&autoplay=1' />"}} />;
      }
    }