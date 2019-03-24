
// -*- JavaScript -*-


class MemberItem extends React.Component {
  render() {
    return (
      <tr>
        <td> {this.props.id}    </td>
        <td> {this.props.firstname} </td>
        <td> {this.props.lastname}  </td>
      </tr>
    );
  }
}

class MemberForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            firstname: '',
            lastname: ''
        };
        this.handleFirstnameChange = this.handleFirstnameChange.bind(this);
        this.handleLastnameChange = this.handleLastnameChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleFirstnameChange(event) {
        console.log('in firstname handler');
        this.setState({firstname: event.target.value});
    }

    handleLastnameChange(event) {
        this.setState({lastname: event.target.value});
    }

    handleSubmit(event) {
        console.info('A new member is added: ' + this.state.firstname + ' ' + this.state.lastname);
        event.preventDefault();
        axios
            .post("/members", null, {
                params: {
                    firstname: this.state.firstname,
                    lastname: this.state.lastname
                }
            })
            .then((result) => {
                console.log('Result of adding member: ' + result);
                this.setState({firstname: '', lastname: ''})
            });
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <input type="text" value={this.state.value} onChange={this.handleFirstnameChange} />
                <input type="text" value={this.state.value} onChange={this.handleLastnameChange} />
                <input type="submit" value="Submit" />
            </form>
        );
    }
}

class MembersList extends React.Component {
  constructor(props) {
    super(props);
    this.state = { members: [] };
  }

  componentDidMount() {
    this.serverRequest =
      axios
        .get("/members")
        .then((result) => {
            console.log('Received members: ' + result.data);
           this.setState({ members: result.data });
        });
  }

  render() {
    const members = this.state.members.map((member, i) => {
      return (
        <MemberItem key={i} id={member.Id} firstname={member.Firstname} lastname={member.Lastname} />
      );
    });

    return (
      <div>
        <table><tbody>
          <tr><th>Id</th><th>Firstname</th><th>Lastname</th></tr>
          {members}
        </tbody></table>
          <MemberForm/>

      </div>
    );
  }
}

class Container extends React.Component {
    render() {
        return (
            <div>
                <MembersList/>
            </div>
        );
    }
}

ReactDOM.render( <Container/>, document.querySelector("#root"));
