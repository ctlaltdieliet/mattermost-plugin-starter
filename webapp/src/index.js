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
            <Icon />,
            () => {
                alert("Hi there");
            },
            "Says hello",
        );
    }
}

window.registerPlugin('be.controlaltdieliet.demoortom.first_mattermost_plugin', new HelloWorldPlugin());
