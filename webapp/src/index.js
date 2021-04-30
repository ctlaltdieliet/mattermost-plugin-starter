import React from 'react';

// Courtesy of https://feathericons.com/
const Icon = () => <i className='icon fa fa-plug'/>;

class HelloWorldPlugin {
    initialize(registry, store) {
        
        registry.registerMainMenuAction(
            'Plugin Menu Item',
            () => console.log("Hi"),
         Icon
        );
        registry.registerChannelHeaderButtonAction(
            // icon - JSX element to use as the button's icon
            <Icon />,
            // action - a function called when the button is clicked, passed the channel and channel member as arguments
            // null,
            () => {
                alert("Hello Tom");
            },
            // dropdown_text - string or JSX element shown for the dropdown button description
            "Hello you",
        );
    }
}

window.registerPlugin('be.controlaltdieliet.demoortom.first_mattermost_plugin', new HelloWorldPlugin());
