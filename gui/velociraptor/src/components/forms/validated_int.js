import "./validated.css";

import React from 'react';
import PropTypes from 'prop-types';

import Form from 'react-bootstrap/Form';

const regexp = new RegExp(`^-?[0-9]*$`);

export default class ValidatedInteger extends React.Component {
    static propTypes = {
        setInvalid: PropTypes.func,
        setValue: PropTypes.func.isRequired,
        value: PropTypes.any,
        placeholder: PropTypes.string,
    };

    state = {
        invalid: false,
    }

    render() {
        let value = this.props.value || "";
        return (
            <>
              <Form.Control placeholder={this.props.placeholder || ""}
                            className={ this.state.invalid && 'invalid' }
                            value={ value }
                            onChange={ (event) => {
                                const newValue = event.target.value;
                                let invalid = true;
                                if (regexp.test(newValue)) {
                                    this.props.setValue(parseInt(newValue));
                                    invalid = false;
                                } else {
                                    invalid = true;
                                }

                                if (this.props.setInvalid) {
                                    this.props.setInvalid(invalid);
                                }
                                this.setState({invalid: invalid});
                            } }
              />
            </>
        );
    }
};
