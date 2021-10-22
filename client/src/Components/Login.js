import React from "react";
import { Button, Card, Input, Form, Modal } from "antd";
import { InfoCircleOutlined } from '@ant-design/icons';
import "./Login.css";

class Login extends React.Component {

  componentDidMount() {
    // todo: send current location to websocket
  }

  copyUrl = (shortenurl) => {
    navigator.clipboard.writeText(shortenurl);
  }

  onFinish = (values) => {
    console.log('Success:', values); 
    fetch('http://localhost:80/api/shorten', {
      method: 'post',
      body: JSON.stringify(values)
    }).then((response) => {
      return response.json();
    }).then((data) => {
      Modal.info({
        content: <div>Shortened URL: <strong>{data.shortenurl}</strong></div>,
        okText: 'Copy',
        onOk: () => this.copyUrl(data.shortenurl)
      });
    }).catch(() => {
      Modal.error({
        content: `Something went wrong!`
      })
    })
  };

  onFinishFailed = (errorInfo) => {
    console.log('Failed:', errorInfo);
  };

  render() {
    return (
      <div className="login-container">
        <Card
          style={{ width: "40rem" }}
          title="Shorten your URL here"
          hoverable
          headStyle={{ fontSize: "20px" }}
          tabProps={{ centered: true }}
        >
          <Form
            name="basic"
            layout="vertical"
            labelCol={{
              span: 8,
            }}
            wrapperCol={{
              span: 25,
            }}
            onFinish={this.onFinish}
            onFinishFailed={this.onFinishFailed}
            autoComplete="off"
            requiredMark={'optional'}
          >
            <Form.Item
              label="URL"
              name="originalurl"
              required tooltip="This is a required field"
              rules={[
                {
                  required: true,
                  message: 'Please input your URL!',
                },
              ]}
            >
              <Input />
            </Form.Item>

            <Form.Item
              label="Shortform"
              tooltip={{ title: 'Provide for custom URL', icon: <InfoCircleOutlined /> }}
              name="shortenurl"
              rules={[
                {
                  required: false,
                },
              ]}
            >
              <Input />
            </Form.Item>

            <Form.Item
              wrapperCol={{
                offset: 0,
                span: 12,
              }}
            >
              <Button type="primary" htmlType="submit">
                Submit
              </Button>
            </Form.Item>
          </Form>
        </Card>
      </div>
    );
  }
}

export default Login;
