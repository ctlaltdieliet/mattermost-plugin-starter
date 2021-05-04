import React from 'react';
import RightSideBar from './components/right_hand_sidebar'
import MyFirstComponent from './components/my_first_component'

const Icon = () => <i className='icon fa fa-plug'/>;
class HelloWorldPlugin {
    initialize(registry, store) {
        registry.registerMainMenuAction(
          "Saying hi",
          () => allert("Hi again"),
        );
        registry.registerChannelHeaderButtonAction(
            // icon - JSX element to use as the button's icon
            <Icon />,
            // action - a function called when the button is clicked, passed the channel and channel member as arguments
            // null,
          
                () => store.dispatch(toggleRHSPlugin),
          
            // dropdown_text - string or JSX element shown for the dropdown button description
            "Says hello",
        );

        const {toggleRHSPlugin} = registry.registerRightHandSidebarComponent(
            RightSideBar,"Applause!");
  
            const {leftsidebarheader}=registry.registerLeftSidebarHeaderComponent(MyFirstComponent)
            
    
    }
}

window.registerPlugin('be.controlaltdieliet.demoortom.first_mattermost_plugin', new HelloWorldPlugin());




