import React from 'react';
import MyFirstComponent from './components/my_first_component'
const Icon = () => <i className='icon fa fa-plug'/>;
class HelloWorldPlugin {
    initialize(registry, store) {
        const {leftsidebarheader}=registry.registerLeftSidebarHeaderComponent(MyFirstComponent)
        registry.registerMainMenuAction(
          "Saying hi",
          () => alert("Hi again"),
        );
        registry.registerChannelHeaderButtonAction(
            // icon - JSX element to use as the button's icon
            <Icon />,
            // action - a function called when the button is clicked, passed the channel and channel member as arguments
            // null,
            () => {
                alert("Hi there");
            },
            // dropdown_text - string or JSX element shown for the dropdown button description
            "Says hello",
        );
    }
}

window.registerPlugin('be.controlaltdieliet.demoortom.first_mattermost_plugin', new HelloWorldPlugin());
